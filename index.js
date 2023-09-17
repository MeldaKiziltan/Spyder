const { Neurosity } = require("@neurosity/sdk");
const robot = require('robotjs');
require("dotenv").config();

const deviceId = process.env.DEVICE_ID || "";
const email = process.env.EMAIL || "";
const password = process.env.PASSWORD || "";

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

    // neurosity.kinesis("moveBackward").subscribe((intent) => {
    //     console.log("Intent: ", intent);
    //     console.log("Backwards!");
    // });

    // subscription.unsubscribe();

    neurosity.kinesis("moveForward", "moveBackward").subscribe((intent) => {
        console.log("Intent: ", intent);
        console.log("Onwards!");
    });

    // sub2.unsubscribe();

    // neurosity.kinesis("rotateLeft").subscribe((intent) => {
    //     console.log("Intent: ", intent);
    //     console.log("Leftie loosey");
    // });


    // neurosity.kinesis("leftArm").subscribe((intent) => {
    //     console.log("Intent: ", intent);
    //     console.log("Left Arm");
    // });

  };
  
  main();

//   const two = async () => {
//     await neurosity
//       .login({
//         email,
//         password
//       })
//       .catch((error) => {
//         console.log(error); 
//         throw new Error(error);
//       });
//     console.log("Logged in");

//     // neurosity.kinesis("moveBackward").subscribe((intent) => {
//     //     console.log("Intent: ", intent);
//     //     console.log("Backwards!");
//     // });

//     neurosity.kinesis("moveForward").subscribe((intent) => {
//         console.log("Intent: ", intent);
//         console.log("Onwards!");
//     });

//     // neurosity.kinesis().pipe()

//     // neurosity.kinesis("rotateLeft").subscribe((intent) => {
//     //     console.log("Intent: ", intent);
//     //     console.log("Leftie loosey");
//     // });


// }

// const three = async () => {
//     await neurosity
//       .login({
//         email,
//         password
//       })
//       .catch((error) => {
//         console.log(error); 
//         throw new Error(error);
//       });
//     console.log("Logged in");

//     // neurosity.kinesis("moveBackward").subscribe((intent) => {
//     //     console.log("Intent: ", intent);
//     //     console.log("Backwards!");
//     // });

//     // neurosity.kinesis("moveForward").subscribe((intent) => {
//     //     console.log("Intent: ", intent);
//     //     console.log("Onwards!");
//     // });

//     neurosity.kinesis("rotateLeft").subscribe((intent) => {
//         console.log("Intent: ", intent);
//         console.log("Leftie loosey");
//     });
// }

// two();
// three();