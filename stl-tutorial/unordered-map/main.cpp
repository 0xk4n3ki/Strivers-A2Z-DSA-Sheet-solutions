#include <bits/stdc++.h>

int main() {
    std::unordered_map<int, int> mp;
    for(int i=1; i<=5; i++) {
        mp.insert({i, i*10});
    }

    std::cout << "Elements present in the map: " << std::endl;
    std::cout << "Key\t element" << std::endl;

    for(auto it=mp.begin(); it!=mp.end(); it++) {
        std::cout << it->first << "\t" << it->second << std::endl;
    }

    int n=2;
    if(mp.find(n) != mp.end()){
        std::cout << n << " is present in map" << std::endl;
    }
    mp.erase(mp.begin());

    std::cout << "Elements after deleting the first element: " << std::endl;
    std::cout << "key\t element" << std::endl;

    for(auto it=mp.begin(); it!=mp.end(); it++){
        std::cout << it->first << "\t" << it->second << std::endl;
    }

    std::cout << "size of the map: " << mp.size() << std::endl;

    if(!mp.empty()) {
        std::cout << "map is not empty" << std::endl;
    }else {
        std::cout << "map is empty" << std::endl;
    }


    auto minElement = std::min_element(mp.begin(), mp.end(), [](const auto &a, const auto &b)
    {
        return a.second < b.second;
    });
    std::cout << "min: " << minElement -> first << " -> " << minElement->second << std::endl;

    mp.clear();
    std::cout << "size of map after clearing all elements: " << mp.size() << std::endl;


    std::vector<int> v {4, 2, 6, 8, 3, 7};
    std::cout << "min: " << *min_element(v.begin(), v.end()) << std::endl;
}