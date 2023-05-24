function calculateProportions(arr) {
  const n = arr.length;
  let positiveCount = 0;
  let negativeCount = 0;
  let zeroCount = 0;

  for (let i = 0; i < n; i++) {
    if (arr[i] > 0) {
      positiveCount++;
    } else if (arr[i] < 0) {
      negativeCount++;
    } else {
      zeroCount++;
    }
  }

  const positiveProportion = (positiveCount / n).toFixed(6);
  const negativeProportion = (negativeCount / n).toFixed(6);
  const zeroProportion = (zeroCount / n).toFixed(6);

  return [positiveProportion, negativeProportion, zeroProportion];
}

// Example usage
const input = "-4 3 -9 0 4 1";
const arr = input.split(" ").map(Number);
const proportions = calculateProportions(arr);

for (let i = 0; i < 3; i++) {
  console.log(proportions[i]);
}