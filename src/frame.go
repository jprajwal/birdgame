package main

import "github.com/jprajwal/birdgame/rendering"

type Frame struct {
	objects []rendering.Renderable
}

func NewFrame() *Frame {
	return &Frame{objects: make([]rendering.Renderable, 0)}
}

func (f *Frame) AddRenderable(r rendering.Renderable) {
	f.objects = append(f.objects, r)
}

func (f *Frame) GetRenderables() []rendering.Renderable {
	return f.Clone().objects
}

func (f *Frame) Clone() *Frame {
	objectsCopy := make([]rendering.Renderable, len(f.objects))
	for i, obj := range f.objects {
		objectsCopy[i] = obj
	}
	return &Frame{objects: objectsCopy}
}

func (f *Frame) Objects() []rendering.Renderable {
	return f.Clone().objects
}
