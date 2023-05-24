const a = "12:01:00PM";
const b = "12:01:00AM";

// Extract time in the format "HH:MM:SS" from string 'a'
const timeA = a.match(/\d{2}:\d{2}:\d{2}/)[0];

// Convert 'timeA' from 12-hour format to 24-hour format
let [hoursA, minutesA, secondsA] = timeA.split(":");
if (a.includes("PM") && hoursA !== "12") {
  hoursA = String(Number(hoursA) + 12);
} else if (a.includes("AM") && hoursA === "12") {
  hoursA = "00";
}

// Extract time in the format "HH:MM:SS" from string 'b'
const timeB = b.match(/\d{2}:\d{2}:\d{2}/)[0];

// Convert 'timeB' from 12-hour format to 24-hour format
let [hoursB, minutesB, secondsB] = timeB.split(":");
if (b.includes("PM") && hoursB !== "12") {
  hoursB = String(Number(hoursB) + 12);
} else if (b.includes("AM") && hoursB === "12") {
  hoursB = "00";
}

// Construct the final output strings
const outputA = `${hoursA}:${minutesA}:${secondsA}`;
const outputB = `${hoursB}:${minutesB}:${secondsB}`;

console.log("Output a:", outputA);  // Output a: 12:01:00
console.log("Output b:", outputB);  // Output b: 00:01:00