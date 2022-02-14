#include <iostream>
#include <queue>
#include <unordered_map>
#include <unordered_set>

using namespace std;

unordered_map<string, vector<string>> m;
queue<string> q;

bool BFS(string name) {
    unordered_set<string> visited;
    while (!q.empty()) {
        string s = q.front();
        q.pop();
        if (visited.count(s)==0) {
            visited.insert(s);
            if (name == s) {
                return true;
            }
            vector<string> v = m[s];
            for (auto i = v.begin(); i!=v.end(); i++) {
                q.push(*i);
            }
        }
    }
    return false;
}

int main(int argc, const char * argv[]) {
    m["you"] = {"Alice", "Claire", "Bob"};
    m["Alice"] = {"Peggy"};
    m["Bob"] = {"Peggy", "Anuj"};
    m["Claire"] = {"Jonny", "Thom"};
    m["Peggy"] = {"you"};
    q.push("you");
    cout<<BFS("Anuj");
    return 0;
}
