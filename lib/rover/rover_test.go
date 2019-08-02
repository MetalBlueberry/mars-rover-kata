package rover_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/metalblueberry/go3d/vec2"
	"github.com/metalblueberry/mars-rover-kata/lib/rover"
	"github.com/metalblueberry/mars-rover-kata/lib/world"
)

var _ = Describe("Rover", func() {
	Describe("When parsing sequence", func() {
		var (
			w *world.Grid
		)
		BeforeEach(func() {
			w = world.New(5, 5)
		})
		It("Should advance 3 steps in X axis", func() {
			sequence := "FFF"

			entity := rover.NewRover(w)
			_, ok := entity.ExecuteSequence(sequence)

			Expect(ok).To(BeTrue())
			Expect(entity.Orientation).To(Equal(vec2.UnitX))
			Expect(entity.Position).To(Equal(vec2.T{3, 0}))
		})
		It("Should advance 3 steps in Y axis", func() {
			sequence := "LFFF"

			entity := rover.NewRover(w)
			_, ok := entity.ExecuteSequence(sequence)

			Expect(ok).To(BeTrue())
			Expect(entity.Orientation).To(Equal(vec2.UnitY))
			Expect(entity.Position).To(Equal(vec2.T{0, 3}))
		})
		It("Should wrap X", func() {
			sequence := "FFFFF"

			entity := rover.NewRover(w)
			_, ok := entity.ExecuteSequence(sequence)

			Expect(ok).To(BeTrue())
			Expect(entity.Orientation).To(Equal(vec2.UnitX))
			Expect(entity.Position).To(Equal(vec2.T{0, 0}))
		})
		It("Should wrap Y", func() {
			sequence := "RFFFFF"

			entity := rover.NewRover(w)
			_, ok := entity.ExecuteSequence(sequence)

			Expect(ok).To(BeTrue())
			Expect(entity.Orientation).To(Equal(vec2.UnitY.Inverted()))
			Expect(entity.Position).To(Equal(vec2.T{0, 0}))
		})
		It("Should walk arround", func() {
			sequence := "FFFFLFFFFLFFFF"

			entity := rover.NewRover(w)
			_, ok := entity.ExecuteSequence(sequence)

			Expect(ok).To(BeTrue())
			Expect(entity.Orientation).To(Equal(vec2.UnitX.Inverted()))
			Expect(entity.Position).To(Equal(vec2.T{0, 4}))
		})
	})
})
