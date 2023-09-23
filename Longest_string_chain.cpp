//https://leetcode.com/problems/longest-string-chain/

class Solution {
public:
    int longestStrChain(vector<string>& a) {
        int  n = a.size();
        auto cmp = [&](string a,string b)->bool{
            return a.size()<b.size();
        };
        sort(a.begin(),a.end(),cmp);
        map<string,int> vis;
        int ans = 0;
        for(int i=0;i<n;i++){
            for(int j=0;j<a[i].size();j++){
                string tem = a[i];
                string gst = tem.erase(j,1);
                vis[a[i]] = max(vis[a[i]], vis[tem]+1);
                ans = max(ans , vis[a[i]]);
            }
        }
        return ans;
    }
};