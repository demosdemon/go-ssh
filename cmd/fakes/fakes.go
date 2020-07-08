package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"sync"

	"github.com/iancoleman/strcase"
	"github.com/maxbrunsfeld/counterfeiter/v6/generator"
	"github.com/sirupsen/logrus"
)

type Interface struct {
	Mode                   generator.FakeMode
	InterfaceName          string
	PackagePath            string
	FakeImplName           string
	DestinationPackageName string
	OutputPath             string
}

func NewInterface(pkg, name string) Interface {
	var fakeImplName, destinationPackageName, outputPath string

	fakeImplName = "Fake" + name
	fileName := strcase.ToSnake(fakeImplName) + ".go"
	if pkg == "." {
		destinationPackageName = "fakes"
		outputPath = "./fakes/" + fileName
	} else {
		pkgName := path.Base(pkg)
		destinationPackageName = pkgName + "_fakes"
		outputPath = "./fakes/" + destinationPackageName + "/" + fileName
	}

	return Interface{
		InterfaceName:          name,
		PackagePath:            pkg,
		FakeImplName:           fakeImplName,
		DestinationPackageName: destinationPackageName,
		OutputPath:             outputPath,
	}
}

func (i *Interface) Logger() *logrus.Entry {
	data, err := json.Marshal(i)
	if err != nil {
		return logrus.WithError(err)
	}
	var fields logrus.Fields
	if err := json.Unmarshal(data, &fields); err != nil {
		return logrus.WithError(err).WithField("data", string(data))
	}
	return logrus.WithFields(fields)
}

var interfaces = []Interface{
	//NewInterface("net", "Addr"),
	NewInterface("net", "Conn"),
	NewInterface("net", "Listener"),
	NewInterface("golang.org/x/crypto/ssh", "Channel"),
	NewInterface("golang.org/x/crypto/ssh", "Conn"),
	//NewInterface("golang.org/x/crypto/ssh", "ConnMetadata"),
	//NewInterface("golang.org/x/crypto/ssh", "NewChannel"),
	NewInterface(".", "Context"),
	NewInterface(".", "Session"),
	//NewInterface(".", "RequestHandler"),
	//NewInterface(".", "ChannelHandler"),
	//NewInterface(".", "Logger"),
	//NewInterface(".", "PublicKey"),
	//NewInterface(".", "Signer"),
}

func main() {
	var wg sync.WaitGroup

	cacher := generator.Cache{}

	pwd, err := os.Getwd()
	if err != nil {
		logrus.WithError(err).Panic("unable to get current working directory")
	}

	for _, iface := range interfaces {
		wg.Add(1)
		iface := iface
		go func() {
			defer wg.Done()

			log := iface.Logger()

			fake, err := generator.NewFake(
				iface.Mode,
				iface.InterfaceName,
				iface.PackagePath,
				iface.FakeImplName,
				iface.DestinationPackageName,
				pwd,
				&cacher,
			)

			if err != nil {
				log.WithError(err).Error("unable to generate faker")
				return
			}

			data, err := fake.Generate(true)
			if err != nil {
				log.WithError(err).Error("unable to generate code")
				return
			}

			dir := path.Dir(iface.OutputPath)
			if err := os.MkdirAll(dir, 0777); err != nil && !os.IsExist(err) {
				log.WithError(err).WithField("dir", dir).Error("unable to create directory")
				return
			}

			if err := ioutil.WriteFile(iface.OutputPath, data, 0666); err != nil {
				log.WithError(err).Error("unable to write file")
				return
			}

			log.Info("file written")
		}()
	}

	wg.Wait()
	logrus.Exit(0)
}
