package logger

import "log/slog"

var censors = map[string]string{}

func CensorReplacer(groups []string, a slog.Attr) (slog.Attr, bool) {
	for k, v := range censors {
		if a.Key == k {
			return slog.String(k, v), true
		}
	}
	return a, false
}
