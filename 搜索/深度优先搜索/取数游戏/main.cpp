#include <iostream>
#include <cstring>
using namespace std;

const int N = 7;
int a[N][N], n, m, t, ans;
int vis[N][N];
int dx[] = {-1, -1, -1, 0, 0, 1, 1, 1};
int dy[] = {-1, 0, 1, -1, 1, -1, 0, 1};

void dfs(int x, int y, int curr) {
    if (y > m) {
        dfs(x + 1, 1, curr);
        return;
    }
    if (x > n) {
        ans = max(ans, curr);
        return;
    }
    dfs(x, y + 1, curr);
    if (!vis[x][y]) {
        vis[x][y]++;
        for (int i = 0; i < 8; i++) {
            int nx = x + dx[i], ny = y + dy[i];
            if (nx >= 1 && nx <= n && ny >= 1 && ny <= m) {
                vis[nx][ny]++;
            }
        }
        dfs(x, y + 1, curr + a[x][y]);
        vis[x][y]--;
        for (int i = 0; i < 8; i++) {
            int nx = x + dx[i], ny = y + dy[i];
            if (nx >= 1 && nx <= n && ny >= 1 && ny <= m) {
                vis[nx][ny]--;
            }
        }
    }
}

int main() {
    cin >> t;
    while (t--) {
        ans = 0;
        memset(vis, 0, sizeof vis);
        cin >> n >> m;
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= m; j++) {
                cin >> a[i][j];
            }
        }
        dfs(1, 1, 0);
        cout << ans << endl;
    }
    return 0;
}