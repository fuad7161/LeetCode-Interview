// https://leetcode.com/problems/reconstruct-itinerary/?envType=daily-question&envId=2023-09-14
class Solution {
public:
    vector<string> findItinerary(vector<vector<string>>& tickets) {
        unordered_map<string,vector<string>> graph;
        for (auto & ticket: tickets){
            graph[ticket[0]].push_back(ticket[1]);
        }
        for(auto &[_,destinations]: graph){
            sort(destinations.rbegin(),destinations.rend());
        }
        vector<string> it;
        function<void(const string&)>dfs = [&](const string& airport){
            while(!graph[airport].empty()){
                string next = graph[airport].back();
                graph[airport].pop_back();
                dfs(next);
            }
            it.push_back(airport);
        };
        dfs("JFK");
        reverse(it.begin(),it.end());
        return it;
    }
};