package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Player struct{
	name string
	guess int
}

type Game struct{
	Player1 Player
	Player2 Player
	rand int
}

func (namwe *Player) playerz(name string){
	namwe.name = name
	
}


func(i *Game)Initialize(){
	i.createGame()
	i.Player1.playerz("Penis")
	i.Player2.playerz("Ass")
	i.gameLoop()
}

func(g *Game)createGame(){
	g.Player1.playerz("")
	g.Player2.playerz("")
	g.rand = rand.Intn(100) + 1
}


func(p *Player)inputGame(){
	var input int
	fmt.Printf("%s Enter your guess: \n", p.name)
	fmt.Scanf("%d", &input)
	p.guess = input
}

func(g *Game)gameLoop(){
	
	for{
		g.Player1.inputGame()
		g.Player2.inputGame()
		p1gap:= math.Abs((float64(g.Player1.guess - g.rand)))
		p2gap:= math.Abs((float64(g.Player2.guess - g.rand)))



		if p1gap<p2gap{
			fmt.Printf("\n%s wins\n",	g.Player1.name)
		}else if p1gap>p2gap{
			fmt.Printf("\n%s wins\n",	g.Player2.name)
		}else{
			fmt.Println("fuck you!")
		}

	}
}

func main(){
	game:= Game{}
	game.Initialize()
}