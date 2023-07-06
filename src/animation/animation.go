package animation

import "time"

// import "fmt"

type Animatable interface {
	MoveHorizontal(delta int)
	MoveVertical(delta int)
}

type Animation interface {
	Done() bool
	Animate(a Animatable)
}

type SlideLeftAnimation struct {
	slide    int
	duration time.Duration
	start    time.Time
	done     bool
	count    int
}

func NewSlideLeftAnimation(slide int, duration int64) *SlideLeftAnimation {
	now := time.Now()
	return &SlideLeftAnimation{
		slide:    slide,
		duration: time.Duration(duration+1) * time.Millisecond,
		start:    now,
		done:     false,
		count:    0,
	}
}

func (a *SlideLeftAnimation) Done() bool {
	a.done = (a.count >= a.slide)
	// fmt.Println("done: ", a.done)
	return a.done
}

func (a *SlideLeftAnimation) Animate(animatable Animatable) {
	// fmt.Println("count: ", a.count)
	if a.Done() {
		return
	}
	interval := time.Duration(a.duration.Milliseconds()/int64(a.slide)) * time.Millisecond
	elapsed := time.Since(a.start)
	intervalCount := elapsed / interval
	pendingCount := int64(intervalCount) - int64(a.count)
	// fmt.Printf("interval: %v, elapsed: %v, intervalCount: %v, pendingCount: %v", interval, elapsed, intervalCount, pendingCount)
	if pendingCount > 0 {
		animatable.MoveHorizontal(int(pendingCount * -1))
		a.count += int(pendingCount)
	}
}
