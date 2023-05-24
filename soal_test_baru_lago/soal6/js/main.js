const a = "07:05:45PM";

// Extract time in the format "HH:MM:SS" from string 'a'
const timeA = a.match(/\d{2}:\d{2}:\d{2}/)[0];

// Convert 'timeA' from 12-hour format to 24-hour format
let [hoursA, minutesA, secondsA] = timeA.split(":");
if (a.includes("PM") && hoursA !== "12") {
  hoursA = String(Number(hoursA) + 12);
} else if (a.includes("AM") && hoursA === "12") {
  hoursA = "00";
}

// Construct the final output string
const outputA = `${hoursA}:${minutesA}:${secondsA}`;

console.log("Output a:", outputA);  // Output a: 19:05:45