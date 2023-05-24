function calculateMinMaxSum(arr) {
  // Sort the array in ascending order
  arr.sort(function(a, b) {
    return a - b;
  });

  // Calculate the sum of the four smallest numbers
  var minSum = arr[0] + arr[1] + arr[2] + arr[3];

  // Calculate the sum of the four largest numbers
  var maxSum = arr[1] + arr[2] + arr[3] + arr[4];

  return minSum + ' ' + maxSum;
}

// Read input from the user
var input = prompt("Enter five space-separated integers:").trim();

// Split the input string into an array of integers
var arr = input.split(' ').map(Number);

// Calculate the minimum and maximum sums
var result = calculateMinMaxSum(arr);

// Print the result
console.log(result);