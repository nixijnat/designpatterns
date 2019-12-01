package state

import "log"

/*
run : -> wait -> stop
hanging -> run -> stop
wait -> hanging -> run -> stop
stop: -> hanging -> run
*/

type state interface {
	setContext(c *context)
	run()
	hanging()
	wait()
	stop()
	name() string
}

type baseState struct {
	ctx *context
}

func (b *baseState) setContext(c *context) {
	b.ctx = c
}

type runStateType struct {
	baseState
}

func (r *runStateType) run()     {}
func (r *runStateType) hanging() {}
func (r *runStateType) wait() {
	r.ctx.setState(r.ctx.lwait)
}
func (r *runStateType) stop() {
	r.ctx.setState(r.ctx.lstop)
}
func (r *runStateType) name() string {
	return "run"
}

type hangingStateType struct {
	baseState
}

func (r *hangingStateType) run() {
	r.ctx.setState(r.ctx.lrun)
}
func (r *hangingStateType) hanging() {}
func (r *hangingStateType) wait()    {}
func (r *hangingStateType) stop() {
	r.ctx.setState(r.ctx.lstop)
}

func (r *hangingStateType) name() string {
	return "hanging"
}

type waitStateType struct {
	baseState
}

func (r *waitStateType) run() {
	r.ctx.setState(r.ctx.lrun)
}
func (r *waitStateType) hanging() {
	r.ctx.setState(r.ctx.lhanging)
}
func (r *waitStateType) wait() {}
func (r *waitStateType) stop() {
	r.ctx.setState(r.ctx.lstop)
}
func (r *waitStateType) name() string {
	return "wait"
}

type stopStateType struct {
	baseState
}

func (r *stopStateType) run() {
	r.ctx.setState(r.ctx.lrun)
}
func (r *stopStateType) hanging() {
	r.ctx.setState(r.ctx.lhanging)
}
func (r *stopStateType) wait() {}
func (r *stopStateType) stop() {}
func (r *stopStateType) name() string {
	return "stop"
}

type context struct {
	lrun     *runStateType
	lhanging *hangingStateType
	lwait    *waitStateType
	lstop    *stopStateType
	cur      state
}

func (c *context) setState(s state) {
	s.setContext(c)
	c.cur = s
}

func (c *context) run() {
	c.cur.run()
}
func (c *context) hanging() {
	c.cur.hanging()
}
func (c *context) wait() {
	c.cur.wait()
}
func (c *context) stop() {
	c.cur.stop()
}

func newContext() *context {
	s := new(stopStateType)
	c := &context{
		lrun:     new(runStateType),
		lhanging: new(hangingStateType),
		lwait:    new(waitStateType),
		lstop:    s,
	}
	c.setState(s)
	return c
}

func example() {
	var ctx = newContext()
	log.Println("stop", ctx.cur.name())
	ctx.hanging()
	log.Println("hanging", ctx.cur.name())
	ctx.wait()
	log.Println("hanging", ctx.cur.name())
	ctx.run()
	log.Println("run", ctx.cur.name())
	ctx.wait()
	log.Println("wait", ctx.cur.name())
	ctx.stop()
	log.Println("stop", ctx.cur.name())
}
