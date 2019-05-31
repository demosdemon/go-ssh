package ssh

import "github.com/sirupsen/logrus"

type Logger interface {
	logrus.Ext1FieldLogger
}

var _ Logger = (*Server)(nil)

func (srv *Server) WithField(key string, value interface{}) *logrus.Entry {
	return srv.logger().WithField(key, value)
}

func (srv *Server) WithFields(fields logrus.Fields) *logrus.Entry {
	return srv.logger().WithFields(fields)
}

func (srv *Server) WithError(err error) *logrus.Entry {
	return srv.logger().WithError(err)
}

func (srv *Server) Tracef(format string, args ...interface{}) {
	srv.logger().Tracef(format, args...)
}

func (srv *Server) Debugf(format string, args ...interface{}) {
	srv.logger().Debugf(format, args...)
}

func (srv *Server) Infof(format string, args ...interface{}) {
	srv.logger().Infof(format, args...)
}

func (srv *Server) Printf(format string, args ...interface{}) {
	srv.logger().Printf(format, args...)
}

func (srv *Server) Warnf(format string, args ...interface{}) {
	srv.logger().Warnf(format, args...)
}

func (srv *Server) Warningf(format string, args ...interface{}) {
	srv.logger().Warningf(format, args...)
}

func (srv *Server) Errorf(format string, args ...interface{}) {
	srv.logger().Errorf(format, args...)
}

func (srv *Server) Fatalf(format string, args ...interface{}) {
	srv.logger().Fatalf(format, args...)
}

func (srv *Server) Panicf(format string, args ...interface{}) {
	srv.logger().Panicf(format, args...)
}

func (srv *Server) Trace(args ...interface{}) {
	srv.logger().Trace(args...)
}

func (srv *Server) Debug(args ...interface{}) {
	srv.logger().Debug(args...)
}

func (srv *Server) Info(args ...interface{}) {
	srv.logger().Info(args...)
}

func (srv *Server) Print(args ...interface{}) {
	srv.logger().Print(args...)
}

func (srv *Server) Warn(args ...interface{}) {
	srv.logger().Warn(args...)
}

func (srv *Server) Warning(args ...interface{}) {
	srv.logger().Warning(args...)
}

func (srv *Server) Error(args ...interface{}) {
	srv.logger().Error(args...)
}

func (srv *Server) Fatal(args ...interface{}) {
	srv.logger().Fatal(args...)
}

func (srv *Server) Panic(args ...interface{}) {
	srv.logger().Panic(args...)
}

func (srv *Server) Traceln(args ...interface{}) {
	srv.logger().Traceln(args...)
}

func (srv *Server) Debugln(args ...interface{}) {
	srv.logger().Debugln(args...)
}

func (srv *Server) Infoln(args ...interface{}) {
	srv.logger().Infoln(args...)
}

func (srv *Server) Println(args ...interface{}) {
	srv.logger().Println(args...)
}

func (srv *Server) Warnln(args ...interface{}) {
	srv.logger().Warnln(args...)
}

func (srv *Server) Warningln(args ...interface{}) {
	srv.logger().Warningln(args...)
}

func (srv *Server) Errorln(args ...interface{}) {
	srv.logger().Errorln(args...)
}

func (srv *Server) Fatalln(args ...interface{}) {
	srv.logger().Fatalln(args...)
}

func (srv *Server) Panicln(args ...interface{}) {
	srv.logger().Panicln(args...)
}
