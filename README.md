# Gen4DP

This is a simple program which generates C++ dynamic programming code from given state-translation equation.

Unit tests are included.

## Example

for the input:

```c
dp[full] -> dp[full - 1] + dp[full - 2] (full >= 2)
      -> 1 (full == 1 or full == 2)
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
	int index, n;
	cin>>n;
	for (index=0; index<n+0; ++index) {
		if (index>=2) {
			dp[index] = dp[index-1]+dp[index-2];
		}
		else if (index==1||index==2) {
			dp[index] = 1;
		}
		else {
			dp[index] = 0;
		}
	}
	cout<<dp[n-1]<<endl;
	return 0;
}
```

Error handling and multi code style support is still working in progress.
