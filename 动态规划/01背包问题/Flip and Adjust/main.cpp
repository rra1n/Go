#include <iostream>
using namespace std;

const int N = 110, M = 1e4 + 10;
int f[N][M], a[N], b[N], n, m;

int main() {
    cin >> n >> m;
    for (int i = 1; i <= n; i++) cin >> a[i] >> b[i];
    f[0][0] = 1;
    for (int i = 1; i <= n; i++) {
        for (int j = 0; j <= m; j++) {
            if (j >= a[i]) f[i][j] |= f[i-1][j-a[i]];
            if (j >= b[i]) f[i][j] |= f[i-1][j-b[i]];
        }
    }
    string ans = "";
    if (f[n][m]) {
        cout << "Yes" << endl;
        for (int i = n; i >= 1; i--) {
            if (m >= a[i] && f[i-1][m-a[i]]) {
                ans = 'H' + ans ;
                m -= a[i];
            } else if (m >= b[i] && f[i-1][m-b[i]]) {
                ans = 'T' + ans;
                m -= b[i];
            }
        }
        cout << ans;
    } else {
        cout << "No" << endl;
    }
    return 0;
}