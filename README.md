# Readme
## Inspiration
Fall lab and design bay cleanout leads to some pretty interesting things being put out at the free tables. In this case, we were drawn in by a motorized Audi Spyder car. And then, we saw the Neurosity Crown headsets, and an idea was born. A single late night call among team members, excited about the possibility of using a kiddy car for something bigger was all it took. Why can't we learn about cool tech and have fun while we're at it?

Spyder is a way we can control cars with our minds. Use cases include remote rescue, non able-bodied individuals, warehouse, and being extremely cool.

## What it does
Spyder uses the Neurosity Crown to take the brainwaves of an individual, train an AI model to detect and identify certain brainwave patterns, and output them as a recognizable output to humans. It's a dry brain-computer interface (BCI) which means electrodes are placed against the scalp to read the brain's electrical activity. By taking advantage of these non-invasive method of reading electrical impulses, this allows for greater accessibility to neural technology.

Collecting these impulses, we are then able to forward these commands to our Viam interface. Viam is a software platform that allows you to easily put together smart machines and robotic projects. It completely changed the way we coded this hackathon. We used it to integrate every single piece of hardware on the car. More about this below! :)

## How we built it
### Mechanical
The manual steering had to be converted to automatic. We did this in SolidWorks by creating a custom 3D printed rack and pinion steering mechanism with a motor mount that was mounted to the existing steering bracket. Custom gear sizing was used for the rack and pinion due to load-bearing constraints. This allows us to command it with a DC motor via Viam and turn the wheel of the car, while maintaining the aesthetics of the steering wheel.

## Hardware
A 12V battery is connected to a custom soldered power distribution board. This powers the car, the boards, and the steering motor. For the DC motors, they are connected to a Cytron motor controller that supplies 10A to both the drive and steering motors via pulse-width modulation (PWM). A custom LED controller and buck converter PCB stepped down the voltage from 12V to 5V for the LED under glow lights and the Raspberry Pi 4. The Raspberry Pi 4 uses the Viam SDK (which controls all peripherals) and connects to the Neurosity Crown for vision software controlling for the motors. All the wiring is custom soldered, and many parts are custom to fit our needs.

## Software
Viam was an integral part of our software development and hardware bringup. It significantly reduced the amount of code, testing, and general pain we'd normally go through creating smart machine or robotics projects. Viam was instrumental in debugging and testing to see if our system was even viable and to quickly check for bugs. The ability to test features without writing drivers or custom code saved us a lot of time. An exciting feature was how we could take code from Viam and merge it with a Go backend which is normally very difficult to do. Being able to integrate with Go was very cool - usually have to do python (flask + SDK). Being able to use Go, we get extra backend benefits without the headache of integration! Additional software that we used was python for the keyboard control client, testing, and validation of mechanical and electrical hardware. We also used JavaScript and node to access the Neurosity Crown, Neurosity SDK and Kinesis API to grab trained AI signals from the console. We then used websockets to port them over to the Raspberry Pi to be used in driving the car.

## Challenges we ran into
Using the Neurosity Crown was the most challenging. Training the AI model to recognize a user's brainwaves and associate them with actions didn't always work. In addition, grabbing this data for more than one action per session was not possible which made controlling the car difficult as we couldn't fully realise our dream. Additionally, it only caught fire once - which we consider to be a personal best. If anything, we created the world's fastest smoke machine.

## Accomplishments that we're proud of
We are proud of being able to complete a full mechatronics system within our 32 hours. We iterated through the engineering design process several times, pivoting multiple times to best suit our hardware availabilities and quickly making decisions to make sure we'd finish everything on time. It's a technically challenging project - diving into learning about neurotechnology and combining it with a new platform - Viam, to create something fun and useful.

## What we learned
Cars are really cool! Turns out we can do more than we thought with a simple kid car. Viam is really cool! We learned through their workshop that we can easily attach peripherals to boards, use and train computer vision models, and even use SLAM! We spend so much time in class writing drivers, interfaces, and code for peripherals in robotics projects, but Viam has it covered. We were really excited to have had the chance to try it out! Neurotech is really cool! Being able to try out technology that normally isn’t available or difficult to acquire and learn something completely new was a great experience.

## What's next for Spyder
- Backflipping car + wheelies
- Fully integrating the Viam CV for human safety concerning reaction time
- Integrating Adhawk glasses and other sensors to help determine user focus and control

## Getting started
With a few simple commands, you too can get started!
You will need to have Node.js and npm installed
You'll need a raspberry pi and the hardware setup, but assuming you have that, next for software, you have:
1. Git clone
`https://github.com/MeldaKiziltan/Spyder.git`
2. npm setup
```
npm install ws
npm install
```
3. Start!
`npm start`