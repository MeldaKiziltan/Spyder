const { Neurosity } = require("@neurosity/sdk");
const robot = require('robotjs');
require("dotenv").config();

// Grabbing environment variables for login and device info
const deviceId = process.env.DEVICE_ID || "";
const email = process.env.EMAIL || "";
const password = process.env.PASSWORD || "";

const webSocket = new WebSocket('ws://10.33.136.125:8081/ws');

const verifyEnvs = (email, password, deviceId) => {
    const invalidEnv = (env) => {
      return env === "" || env === 0;
    };
    if (invalidEnv(email) || invalidEnv(password) || invalidEnv(deviceId)) {
      console.error(
        "Please verify deviceId, email and password are in .env file, quitting..."
      );
      process.exit(0);
    }
  };

  verifyEnvs(email, password, deviceId);
  
  console.log(`${email} attempting to authenticate to ${deviceId}`);

  const neurosity = new Neurosity({
    deviceId
  });

  const main = async () => {
    await neurosity
      .login({
        email,
        password
      })
      .catch((error) => {
        console.log(error); 
        throw new Error(error);
      });
    console.log("Logged in");

    // Trained actions via Kinesis API and Neurosity SDK
    neurosity.kinesis("moveForward").subscribe((intent) => {
        console.log("Intent: ", intent);
        webSocket.send('drive,0.3')
        console.log("Onwards!");
    });

    // neurosity.kinesis("moveBackward").subscribe((intent) => {
    //     console.log("Intent: ", intent);
    //     console.log("Backwards!");
    // });


    // neurosity.kinesis("rotateLeft").subscribe((intent) => {
    //     console.log("Intent: ", intent);
    //     console.log("Lefty loosey");
    // });

  };
  
  main();