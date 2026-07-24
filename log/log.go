package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

type Logger struct {
	Out     io.Writer
	Webhook string
	MinLvl  Level
}

func New(w io.Writer, webhook string, lvl Level) *Logger {
	if w == nil {
		w = os.Stdout
	}
	return &Logger{
		Out:     w,
		Webhook: webhook,
		MinLvl:  lvl,
	}
}

func (l *Logger) Log(lvl Level, msg string, fields map[string]interface{}) {
	if lvl < l.MinLvl {
		return
	}

	entry := make(map[string]interface{})
	entry["ts"] = time.Now().UTC().Format(time.RFC3339)
	entry["level"] = lvlString(lvl)
	entry["msg"] = msg

	for k, v := range fields {
		entry[k] = v
	}

	buf, _ := json.Marshal(entry)
	fmt.Fprintln(l.Out, string(buf))

	if lvl >= ERROR && l.Webhook != "" {
		go l.sendDiscord(msg, fields)
	}
}

func (l *Logger) Info(msg string, fields map[string]interface{})  { l.Log(INFO, msg, fields) }
func (l *Logger) Warn(msg string, fields map[string]interface{})  { l.Log(WARN, msg, fields) }
func (l *Logger) Error(msg string, fields map[string]interface{}) { l.Log(ERROR, msg, fields) }

func (l *Logger) sendDiscord(msg string, fields map[string]interface{}) {
	type Field struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	type Embed struct {
		Title  string  `json:"title"`
		Color  int     `json:"color"`
		Fields []Field `json:"fields"`
	}
	type Payload struct {
		Embeds []Embed `json:"embeds"`
	}

	var f []Field
	for k, v := range fields {
		f = append(f, Field{Name: k, Value: fmt.Sprintf("%v", v)})
	}

	pld := Payload{
		Embeds: []Embed{
			{
				Title:  "🔴 " + msg,
				Color:  15158332,
				Fields: f,
			},
		},
	}

	b, _ := json.Marshal(pld)
	c := &http.Client{Timeout: 5 * time.Second}
	_, _ = c.Post(l.Webhook, "application/json", bytes.NewBuffer(b))
}

func lvlString(l Level) string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}
