/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co.,Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/08/01        Liu JiaChang
 */

package log

import (
	"go.uber.org/zap"
)

// RecordLog is receiver.
type RecordLog struct{}

var (
	// Logger common logger
	Logger *RecordLog
	zapLog *zap.Logger
	sugar  *zap.SugaredLogger
)

func init() {
	Logger = &RecordLog{}
	zapLog, _ = zap.NewDevelopment()
	sugar = zapLog.Sugar()
}

func (l *RecordLog) Error(desc string, err error) {
	zapLog.Error(desc, zap.Error(err))
}

// Debug debug log
func (l *RecordLog) Debug(a ...interface{}) {
	sugar.Debug(a)
}

// Fatal shutdown if fail
func (l *RecordLog) Fatal(a ...interface{}) {
	sugar.Fatal(a)
}

// Info look some information
func (l *RecordLog) Info(a ...interface{}) {
	sugar.Info(a)
}

// Warn warn
func (l *RecordLog) Warn(a ...interface{}) {
	sugar.Warn(a)
}
