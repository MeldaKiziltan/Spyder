import keyboard
import time
import asyncio
from websockets.sync.client import connect

with connect("ws://10.33.136.125:8081/ws") as websocket:
    while True:
        if keyboard.is_pressed("n"):
            websocket.send("drive,1.0")
            print("w")
        elif keyboard.is_pressed("w"):
            websocket.send("drive,0.5")
            print("w")
        elif keyboard.is_pressed("s"):
            websocket.send("drive,-0.5")
            print("s")
        else:
            # send no drive
            websocket.send("drive,stop")

        if keyboard.is_pressed("a"):
            websocket.send("steering,-0.20")
            print("a")
            time.sleep(0.1)
        elif keyboard.is_pressed("d"):
            websocket.send("steering,0.20")
            print("d")
        else: 
            # send no steering changes
            websocket.send("steering,stop")
        time.sleep(0.1)
