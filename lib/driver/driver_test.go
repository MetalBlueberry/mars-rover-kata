package driver_test

import (
	"fmt"
	"io"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/metalblueberry/go3d/vec2"
	"github.com/metalblueberry/mars-rover-kata/lib/driver"
	"github.com/metalblueberry/mars-rover-kata/lib/rover"
	"github.com/metalblueberry/mars-rover-kata/lib/world"
)

func DrawWorld(out io.Writer, Grid *world.Grid, Driver *driver.Driver) {

	maxX, maxY := Grid.Limits()

	fmt.Fprint(out, " +")
	for x := int64(0); x < maxX; x++ {
		if x%5 == 0 {
			fmt.Fprint(out, "v")
		} else {
			fmt.Fprint(out, "_")
		}
	}
	fmt.Fprint(out, "+")
	fmt.Fprint(out, "\n")
	for y := int64(0); y < maxY; y++ {
		fmt.Fprint(out, "⎹ ")
		for x := int64(0); x < maxX; x++ {
			if i, j := Driver.Position(); x == i && y == j {
				fmt.Fprint(out, "O")
				continue
			}
			if Grid.IsBlocked(x, y) {
				fmt.Fprint(out, "X")
				continue
			}
			fmt.Fprint(out, "-")
		}
		fmt.Fprint(out, "⎸")
		fmt.Fprint(out, "\n")
	}
	fmt.Fprint(out, " +")
	for x := int64(0); x < maxX; x++ {
		if x%5 == 0 {
			fmt.Fprint(out, "^")
		} else {
			fmt.Fprint(out, "‾")
		}
	}
	fmt.Fprint(out, "+")
	fmt.Fprint(out, "\n")

}

var _ = Describe("Driver", func() {
	var (
		World  *world.Grid
		Rover  *rover.Rover
		Driver *driver.Driver
	)
	BeforeEach(func() {
		World = world.New(11, 5)
		Rover = rover.NewRover()
		Driver = driver.NewDriver(vec2.T{5, 2}, driver.North, World, Rover)
	})
	JustBeforeEach(func() {
		DrawWorld(GinkgoWriter, World, Driver)
	})
	AfterEach(func() {
		DrawWorld(GinkgoWriter, World, Driver)
	})
	Describe("When parsing sequence", func() {

		It("Should advance 3 steps in Y axis", func() {
			sequence := "FFF"

			_, ok := Driver.ExecuteSequence(sequence)

			Expect(ok).To(BeTrue())
			Expect(Rover.Orientation).To(Equal(driver.North.Vector()))
			Expect(Rover.Position).To(Equal(vec2.T{5, 4}))
		})
		It("Should advance 3 steps in Y axis", func() {
			sequence := "LFFF"

			_, ok := Driver.ExecuteSequence(sequence)

			Expect(ok).To(BeTrue())
			Expect(Rover.Orientation).To(Equal(driver.East.Vector()))
			Expect(Rover.Position).To(Equal(vec2.T{8, 2}))
		})
		Describe("In negative direction", func() {
			It("Should wrap Y ", func() {
				sequence := "FFFFF"

				_, ok := Driver.ExecuteSequence(sequence)

				Expect(ok).To(BeTrue())
				Expect(Rover.Orientation).To(Equal(driver.North.Vector()))
				Expect(Rover.Position).To(Equal(vec2.T{5, 2}))
			})
			It("Should wrap X ", func() {
				sequence := "RFFFFFF"

				_, ok := Driver.ExecuteSequence(sequence)

				Expect(ok).To(BeTrue())
				Expect(Rover.Orientation).To(Equal(driver.Weast.Vector()))
				Expect(Rover.Position).To(Equal(vec2.T{10, 2}))
			})
		})
		Describe("In Positive direction", func() {
			It("Should wrap Y ", func() {
				sequence := "RRFFFFF"

				_, ok := Driver.ExecuteSequence(sequence)

				Expect(ok).To(BeTrue())
				Expect(Rover.Orientation).To(Equal(driver.South.Vector()))
				Expect(Rover.Position).To(Equal(vec2.T{5, 2}))
			})
			It("Should wrap X ", func() {
				sequence := "LFFFFFF"

				_, ok := Driver.ExecuteSequence(sequence)

				Expect(ok).To(BeTrue())
				Expect(Rover.Orientation).To(Equal(driver.East.Vector()))
				Expect(Rover.Position).To(Equal(vec2.T{0, 2}))
			})
		})
		It("Should walk arround and return to position L", func() {
			sequence := "FFFLFFFLFFFLFFF"

			_, ok := Driver.ExecuteSequence(sequence)

			Expect(ok).To(BeTrue())
			Expect(Rover.Orientation).To(Equal(driver.Weast.Vector()))
			Expect(Rover.Position).To(Equal(vec2.T{5, 2}))
		})
		It("Should walk arround and return to position R", func() {
			sequence := "FFFRFFFRFFFRFFF"

			_, ok := Driver.ExecuteSequence(sequence)

			Expect(ok).To(BeTrue())
			Expect(Rover.Orientation).To(Equal(driver.East.Vector()))
			Expect(Rover.Position).To(Equal(vec2.T{5, 2}))
		})
	})
	Describe("When finding obstacles", func() {
		BeforeEach(func() {
			World.Get(3, 2).Blocked = true
			World.Get(7, 2).Blocked = true
			World.Get(5, 0).Blocked = true
			World.Get(5, 4).Blocked = true
		})
		It("Should stop if colides going Noth", func() {
			sequence := "FFF"

			stop, ok := Driver.ExecuteSequence(sequence)

			Expect(ok).To(BeFalse())
			Expect(stop).To(Equal(1))
			Expect(Rover.Orientation).To(Equal(driver.North.Vector()))
			Expect(Rover.Position).To(Equal(vec2.T{5, 1}))

		})
		It("Should stop if colides going South", func() {
			sequence := "RRFFF"

			stop, ok := Driver.ExecuteSequence(sequence)

			Expect(ok).To(BeFalse())
			Expect(stop).To(Equal(3))
			Expect(Rover.Orientation).To(Equal(driver.South.Vector()))
			Expect(Rover.Position).To(Equal(vec2.T{5, 3}))

		})
		It("Should stop if colides going Weast", func() {
			sequence := "RFFF"

			stop, ok := Driver.ExecuteSequence(sequence)

			Expect(ok).To(BeFalse())
			Expect(stop).To(Equal(2))
			Expect(Rover.Orientation).To(Equal(driver.Weast.Vector()))
			Expect(Rover.Position).To(Equal(vec2.T{4, 2}))

		})
		It("Should stop if colides going East", func() {
			sequence := "LFFF"

			stop, ok := Driver.ExecuteSequence(sequence)

			Expect(ok).To(BeFalse())
			Expect(stop).To(Equal(2))
			Expect(Rover.Orientation).To(Equal(driver.East.Vector()))
			Expect(Rover.Position).To(Equal(vec2.T{6, 2}))

		})
	})
})
