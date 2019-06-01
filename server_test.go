package ssh_test

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	. "github.com/demosdemon/go-ssh"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"net"
	"regexp"
	"strings"
	"testing"
	"text/template"
)

func init() {
	var err error

	hostSigner, err = generateSigner(privateKeyBits)
	if err != nil {
		panic(err)
	}

	clientSigner, err = generateSigner(privateKeyBits)
	if err != nil {
		panic(err)
	}

	password, err = generatePassword(128)
	if err != nil {
		panic(err)
	}
}

const (
	privateKeyBits = 2048

	genericMessage = `Hello {{ .User }}!
Command:{{ range .Command }} {{ quote . }}{{ end }}
Environment:
{{- range .Environ }}
  - {{ . }}
{{- end }}

{{ replicate "*" 80 }}
{{ center "This is a test server. Sessions are not permitted." 80 }}
{{ replicate "*" 80 }}
`
)

var (
	hostSigner   Signer
	clientSigner Signer
	password     string

	templateFuncs = template.FuncMap{
		"quote":     quote,
		"replicate": replicate,
		"center":    center,
	}

	genericMessageTemplate = template.Must(template.New("").Funcs(templateFuncs).Parse(genericMessage))

	unsafeRE = regexp.MustCompile(`[^\w@%+=:,./-]`)
)

func quote(s string) string {
	if s == "" {
		return "''"
	}

	if !unsafeRE.MatchString(s) {
		return s
	}

	return fmt.Sprintf("'%s'", strings.ReplaceAll(s, "'", `'"'"'`))
}

func replicate(s string, times int) string {
	var b strings.Builder
	for idx := 0; idx < times; idx++ {
		b.WriteString(s)
	}
	return b.String()
}

func center(s string, width int) string {
	s = strings.TrimSpace(s)

	padding := width - len(s)
	if padding <= 0 {
		return s
	}

	lpad := padding / 2
	s = replicate(" ", lpad) + s
	rpad := width - len(s)
	return s + replicate(" ", rpad)
}

func generateSigner(bits int) (Signer, error) {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}

	signer, err := NewSignerFromSigner(key)
	if err != nil {
		return nil, err
	}

	return signer, nil
}

func generatePassword(n int) (string, error) {
	buf := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		return "", err
	}

	return base64.RawStdEncoding.EncodeToString(buf[:]), nil
}

func genericHandler(sess Session) {
	err := genericMessageTemplate.Execute(sess, sess)
	if err != nil {
		_, _ = fmt.Fprintf(sess.Stderr(), "\nerror executing template: %v\n", err)
	}
}

func withKeyboardInteractiveLogin(srv *Server) error {
	srv.KeyboardInteractiveHandler = keyboardInteractiveLogin
	return nil
}

func keyboardInteractiveLogin(ctx Context, challenger ssh.KeyboardInteractiveChallenge) bool {
	user := ctx.User()
	instruction := "You know what to do!"
	questions := []string{
		"What is the airspeed velocity of an unladen swallow?",
	}
	echos := []bool{true}

	answers, err := challenger(user, instruction, questions, echos)
	if err != nil {
		ctx.Logger().WithError(err).Warn("error in keyboard interactive response")
		return false
	}

	if len(answers) < 1 {
		ctx.Logger().Warn("expected at least one answer, got 0")
		return false
	}

	if strings.EqualFold(answers[0], "What do you mean? African or European swallow?") {
		ctx.Logger().Info("this person is a Monty Python fan")
		return true
	}

	if strings.EqualFold(answers[0], "Roughly 11 meters per second, or 24 miles an hour.") {
		ctx.Logger().Info("this person is a nerd (or very literal)")
		return true
	}

	questions = []string{"What is your password?"}
	echos = []bool{false}

	answers, err = challenger(user, instruction, questions, echos)
	if err != nil {
		ctx.Logger().WithError(err).Warn("error in keyboard interactive response")
		return false
	}

	if len(answers) < 1 {
		ctx.Logger().Warn("expected at least one answer, got 0")
		return false
	}

	if answers[0] == password {
		ctx.Logger().Info("this person knows our secrets")
		return true
	}

	ctx.Logger().Info("this person is a fraud")
	return false
}

func withPasswordLogin(srv *Server) error {
	srv.PasswordHandler = passwordLogin
	return nil
}

func passwordLogin(_ Context, pw string) bool {
	return password == pw
}

func withPublicKeyLogin(srv *Server) error {
	srv.PublicKeyHandler = publicKeyLogin
	return nil
}

func publicKeyLogin(_ Context, pk PublicKey) bool {
	return KeysEqual(clientSigner.PublicKey(), pk)
}

func server(handler Handler, logger Logger, options ...Option) (*Server, net.Addr, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, nil, err
	}

	log := logger
	if log == nil {
		log = logrus.StandardLogger()
	}
	log = log.WithField("addr", l.Addr().String())

	srv := &Server{
		Addr:        l.Addr().String(),
		Handler:     handler,
		HostSigners: []Signer{hostSigner},
		Version:     "ssh-test",
		Logger:      logger,
	}

	for _, opt := range options {
		err := srv.SetOption(opt)
		if err != nil {
			_ = l.Close()
			return nil, nil, err
		}
	}

	log.Info("ssh server listening")
	go func() {
		_ = srv.Serve(l)
	}()

	return srv, l.Addr(), nil
}

func client(addr net.Addr, auth ...ssh.AuthMethod) (*ssh.Client, error) {
	cfg := ssh.ClientConfig{
		User:            "random test user",
		Auth:            auth,
		HostKeyCallback: ssh.FixedHostKey(hostSigner.PublicKey()),
		ClientVersion:   "SSH-2.0-ssh-test-client",
	}

	return ssh.Dial(addr.Network(), addr.String(), &cfg)
}

func TestServer_NoAuth(t *testing.T) {
	logger := logrus.Logger{
		Out:       ioutil.Discard,
		Hooks:     make(logrus.LevelHooks),
		Formatter: &testFormatter,
		Level:     logrus.TraceLevel,
		ExitFunc:  func(code int) { panic(code) },
	}

	hook := test.NewLocal(&logger)

	srv, addr, err := server(genericHandler, &logger, NoPty())
	require.NoError(t, err, "error opening server")

	cl, err := client(addr)
	require.NoError(t, err, "error opening client")

	sess, err := cl.NewSession()
	require.NoError(t, err, "error opening session")

	var stdout, stderr bytes.Buffer

	sess.Stdout = &stdout
	sess.Stderr = &stderr

	require.NoError(t, sess.Start("echo test '123 456'"), "error starting command")
	require.NoError(t, sess.Wait(), "error waiting for session")
	require.EqualError(t, sess.Close(), "EOF")
	require.NoError(t, cl.Close())
	require.NoError(t, srv.Shutdown(context.Background()))

	assert.Equal(t, stdout.String(), `Hello random test user!
Command: echo test '123 456'
Environment:

********************************************************************************
               This is a test server. Sessions are not permitted.               
********************************************************************************
`)
	assert.Equal(t, stderr.String(), "")
	assert.Equal(t, 3, len(hook.AllEntries()))
}
