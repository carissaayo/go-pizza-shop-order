package main

type NotificationManager struct{
	clients map[string]map[chan string]bool
}

clients map[string]map[chan string]bool={
	"order:123":{
		0xc0000a4000: true,
		0xc0000a4060:true,
	},
	"order:456":{
		0xc0000a40c0:true,
	},
	"admin:new_orders":{
		0xc0000a4180:true,
	},
}