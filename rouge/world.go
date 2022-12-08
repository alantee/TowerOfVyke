package main

import (
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var position *ecs.Component
var renderable *ecs.Component
var monster *ecs.Component
var health *ecs.Component
var meleeWeapon *ecs.Component
var armor *ecs.Component
var name *ecs.Component
var userMessage *ecs.Component

func InitializeWorld(startingLevel Level) (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	player := manager.NewComponent()
	position = manager.NewComponent()
	renderable = manager.NewComponent()
	movable := manager.NewComponent()
	monster = manager.NewComponent()
	health = manager.NewComponent()
	meleeWeapon = manager.NewComponent()
	armor = manager.NewComponent()
	name = manager.NewComponent()
	userMessage = manager.NewComponent()

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}
	redBatImg, _, err := ebitenutil.NewImageFromFile("assets/red_bat.png")
	if err != nil {
		log.Fatal(err)
	}
	shadollImg, _, err := ebitenutil.NewImageFromFile("assets/shadoll.png")
	if err != nil {
		log.Fatal(err)
	}
	orcImg, _, err := ebitenutil.NewImageFromFile("assets/orc.png")
	if err != nil {
		log.Fatal(err)
	}

	//Get First Room
	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()

	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, &Renderable{
			Image: playerImg,
		}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: x,
			Y: y,
		}).
		AddComponent(health, &Health{
			MaxHealth:     100,
			CurrentHealth: 100,
		}).
		AddComponent(meleeWeapon, &MeleeWeapon{
			Name:          "Daggers of the Axum",
			MinimumDamage: 9,
			MaximumDamage: 15,
			ToHitBonus:    3,
		}).
		AddComponent(armor, &Armor{
			Name:       "Iron Armor",
			Defense:    8,
			ArmorClass: 11,
		}).
		AddComponent(name, &Name{Label: "Prince Kwame"}).
		AddComponent(userMessage, &UserMessage{
			AttackMessage:    "",
			DeadMessage:      "",
			GameStateMessage: "",
		})

	//Add a Monster in each room except the player's room
	for _, room := range startingLevel.Rooms {
		if room.X1 != startingRoom.X1 {
			mX, mY := room.Center()

			mobSpawn := GetDiceRoll(3)

			switch mobSpawn {
			case 1:
				manager.NewEntity().
					AddComponent(monster, &Monster{}).
					AddComponent(renderable, &Renderable{
						Image: redBatImg,
					}).
					AddComponent(position, &Position{
						X: mX,
						Y: mY,
					}).
					AddComponent(health, &Health{
						MaxHealth:     25,
						CurrentHealth: 25,
					}).
					AddComponent(meleeWeapon, &MeleeWeapon{
						Name:          "Wings",
						MinimumDamage: 6,
						MaximumDamage: 12,
						ToHitBonus:    4,
					}).
					AddComponent(armor, &Armor{
						Name:       "Skin",
						Defense:    4,
						ArmorClass: 7,
					}).
					AddComponent(name, &Name{Label: "VexBat"}).
					AddComponent(userMessage, &UserMessage{
						AttackMessage:    "",
						DeadMessage:      "",
						GameStateMessage: "",
					})
			case 2:
				manager.NewEntity().
					AddComponent(monster, &Monster{}).
					AddComponent(renderable, &Renderable{
						Image: orcImg,
					}).
					AddComponent(position, &Position{
						X: mX,
						Y: mY,
					}).
					AddComponent(health, &Health{
						MaxHealth:     35,
						CurrentHealth: 35,
					}).
					AddComponent(meleeWeapon, &MeleeWeapon{
						Name:          "Machete",
						MinimumDamage: 6,
						MaximumDamage: 15,
						ToHitBonus:    3,
					}).
					AddComponent(armor, &Armor{
						Name:       "Leather",
						Defense:    6,
						ArmorClass: 9,
					}).
					AddComponent(name, &Name{Label: "Meta-Orc"}).
					AddComponent(userMessage, &UserMessage{
						AttackMessage:    "",
						DeadMessage:      "",
						GameStateMessage: "",
					})
			case 3:
				manager.NewEntity().
					AddComponent(monster, &Monster{}).
					AddComponent(renderable, &Renderable{
						Image: shadollImg,
					}).
					AddComponent(position, &Position{
						X: mX,
						Y: mY,
					}).
					AddComponent(health, &Health{
						MaxHealth:     30,
						CurrentHealth: 30,
					}).
					AddComponent(meleeWeapon, &MeleeWeapon{
						Name:          "Shadow Blade",
						MinimumDamage: 12,
						MaximumDamage: 16,
						ToHitBonus:    6,
					}).
					AddComponent(armor, &Armor{
						Name:       "Cape of the Wretched",
						Defense:    8,
						ArmorClass: 9,
					}).
					AddComponent(name, &Name{Label: "Shadoll Warrior"}).
					AddComponent(userMessage, &UserMessage{
						AttackMessage:    "",
						DeadMessage:      "",
						GameStateMessage: "",
					})
			}
		}
	}

	players := ecs.BuildTag(player, position, health, meleeWeapon, armor, name, userMessage)
	tags["players"] = players

	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables

	monsters := ecs.BuildTag(monster, position, health, meleeWeapon, armor, name, userMessage)
	tags["monsters"] = monsters

	messengers := ecs.BuildTag(userMessage)
	tags["messengers"] = messengers

	return manager, tags
}
