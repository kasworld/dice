// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package dice

import (
	"fmt"
	"math/rand"
	"sort"
)

type Dice struct {
	count   uint8
	side    uint8
	adj     int8
	discard uint8
}

func NewDiceDiscardN(count int, side int, adj int, discard int) *Dice {
	dice := Dice{
		count:   uint8(count),
		side:    uint8(side),
		adj:     int8(adj),
		discard: uint8(discard),
	}
	return &dice
}
func NewDiceN(count int, side int, adj int) *Dice {
	return NewDiceDiscardN(count, side, adj, 0)
}

func NewDice(side int) *Dice {
	return NewDiceDiscardN(1, side, 0, 0)
}

func (d *Dice) Roll() int {
	rtn := make([]int, d.count)
	for i := uint8(0); i < d.count; i++ {
		rtn[i] = rand.Intn(int(d.side)) + 1
	}
	if d.discard > 0 {
		sort.Ints(rtn)
	}
	sum := 0
	for i := d.discard; i < d.count; i++ {
		sum += rtn[i]
	}
	return sum + int(d.adj)
}

func (d Dice) String() string {
	if d.discard > 0 {
		return fmt.Sprintf("(%v-%v)d%v+%v", d.count, d.discard, d.side, d.adj)
	} else {
		if d.adj != 0 {
			return fmt.Sprintf("%vd%v+%v", d.count, d.side, d.adj)
		} else {
			return fmt.Sprintf("%vd%v", d.count, d.side)
		}
	}
}
