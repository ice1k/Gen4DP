/**
sample:

dp[i] -> dp[i - 1] + dp[i - 2] (i >= 2)
      -> 1 (i == 1 || i == 2)
      -> 0 (else)
*/

package dp
