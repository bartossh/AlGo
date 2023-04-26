package crackrsa

// To create custom (e, n, d) values you may use https://www.mobilefish.com/services/rsa_key_generation/rsa_key_generation.php

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessCrack(t *testing.T) {
	n, ok := (&big.Int{}).SetString("24051723933323373230335109652699872887260372863633030520380856590934224554506308944154529656903683098544282868895265857723676740447085769973038138116162852753658181861191950778361549639563565516085451073539560657386103501608592321148669427604194877552133864887585897064910317370632491325912646759075452895764136071794899761625652745642888012193592843601786282707419064157922868466879644136792854722277212465067471658496818060980989808791352963906077940588038623347540668963885547785982543883250789113853569537794783330309654648546163063571756203834919697878945651911998161025323667873893944714006021586935213636888431", 10)
	assert.True(t, ok)
	d, ok := (&big.Int{}).SetString("20859605057389981400415296665239606253551311979432043299936333792698939369418558891569637169366135826146428643134992692481438916188899523620207130817470747633629513081286743218201811495234043370443885950972963184234382668232155560092302387896834347699555010854105235260577040893379009940545782216749159515118484219566373157731404293321389017417036945992984437162056145246504943473128453889715274064071687926343900718250671226003207988553491071490774949729393790264296526140962891140650428560103645538027632465103573248308915991466476312603275778085679414182339076676621372222055380237829179961993191380693342799887257", 10)
	assert.True(t, ok)

	testCases := []struct {
		n []byte
		e []byte
		d []byte
	}{
		{
			n: big.NewInt(63648259).Bytes(),
			e: big.NewInt(65537).Bytes(),
			d: big.NewInt(27903761).Bytes(),
		},
		{
			n: n.Bytes(),
			e: big.NewInt(65537).Bytes(),
			d: d.Bytes(),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			d, err := CrackRSA(tc.e, tc.n)
			assert.Nil(t, err)
			assert.Equal(t, tc.d, d)
		})
	}
}