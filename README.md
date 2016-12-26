# Gen4DP

This is a simple program which generates C++ dynamic programming code from given state-translation equation.

## Example

for the input:

```c
dp[i] -> dp[i - 1] + dp[i - 2] (i >= 2)
      -> 1 (i == 1 or i == 2)
      -> 0 (else)
```

It generates:

```c
#include <iostream>
#include <algorithm>
using namespace std;
#define SIZE 10001
#define number int
number dp[SIZE];

int main(const int argc, const char *argv[]) {
	int i, n;
	cin>>n;
	for (i=0; i<n+0; ++i) {
		if (i>=2) {
			dp[i] = dp[i-1]+dp[i-2];
		}
		else if (i==1||i==2) {
			dp[i] = 1;
		}
		else {
			dp[i] = 0;
		}
	}
	cout<<dp[n-1]<<endl;
	return 0;
}
```

Error handling and multi-code style support is still working in progress.
