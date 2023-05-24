function miniMaxSum(arr) {
  // Sort the array in ascending order
  arr.sort((a, b) => a - b);

  // Calculate the minimum sum by summing the first four elements
  var minSum = arr.slice(0, 4).reduce((a, b) => a + b, 0);

  // Calculate the maximum sum by summing the last four elements
  var maxSum = arr.slice(1).reduce((a, b) => a + b, 0);

  // Print the minimum and maximum sums as a single line
  console.log(minSum + " " + maxSum);
}

// Example usage
var arr = [1, 3, 5, 7, 9];
miniMaxSum(arr);