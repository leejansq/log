package log

import "github.com/sirupsen/logrus"

type Entry struct {
	*logrus.Entry
}

func (l *Entry) Child() *Entry {
	entry := *l.Entry
	return &Entry{&entry}
}

func New() *Entry {
	return &Entry{logrus.NewEntry(logrus.New())}
}
