#include <iostream>
#include <vector>
using namespace std;

// 定义方向数组（左上右下）
const int dx[4] = {0, -1, 0, 1};
const int dy[4] = {-1, 0, 1, 0};

// DFS 函数
void dfs(vector<vector<int>>& maze, vector<vector<bool>>& visited, vector<pair<int, int>>& path, pair<int, int> start, pair<int, int> end, int m, int n) {
    int x = start.first;
    int y = start.second;

    // 检查边界条件
    if (x < 0 || x >= m || y < 0 || y >= n || maze[x][y] == 0 || visited[x][y]) {
        return;
    }

    // 标记当前节点为已访问
    visited[x][y] = true;
    path.push_back({x, y});

    // 如果到达终点，输出路径
    if (x == end.first && y == end.second) {
        for (size_t i = 0; i < path.size(); ++i) {
            if (i > 0) cout << "->";
            cout << "(" << path[i].first + 1 << "," << path[i].second + 1 << ")";
        }
        cout << endl;
    } else {
        // 按照优先顺序（左上右下）递归搜索
        for (int i = 0; i < 4; ++i) {
            int nx = x + dx[i];
            int ny = y + dy[i];
            dfs(maze, visited, path, {nx, ny}, end, m, n);
        }
    }

    // 回溯
    path.pop_back();
    visited[x][y] = false;
}

int main() {
    int m, n;
    cin >> m >> n;

    // 读取迷宫
    vector<vector<int>> maze(m, vector<int>(n));
    for (int i = 0; i < m; ++i) {
        for (int j = 0; j < n; ++j) {
            cin >> maze[i][j];
        }
    }

    // 读取起点和终点
    pair<int, int> start, end;  
    cin >> start.first >> start.second;
    cin >> end.first >> end.second;

    // 转换为 0-based 坐标
    start.first -= 1;
    start.second -= 1;
    end.first -= 1;
    end.second -= 1;

    // 初始化访问标记和路径
    vector<vector<bool>> visited(m, vector<bool>(n, false));
    vector<pair<int, int>> path;

    // 开始 DFS
    dfs(maze, visited, path, start, end, m, n);

    // 如果没有找到任何路径，输出 -1
    if (path.empty()) {
        cout << -1 << endl;
    }

    return 0;
}