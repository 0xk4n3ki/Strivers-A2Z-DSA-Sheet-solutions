#include <bits/stdc++.h>

int main() {
    std::map<int, int> mp;
    for(int i=1; i<=5;i++){
        mp.insert({i, i*10});
    }

    std::cout << "elements present in the map: " << std::endl;
    std::cout << "key\t element" << std::endl;
    for(auto it=mp.begin(); it!=mp.end(); it++) {
        std::cout << it->first << "\t" << it->second << std::endl;
    }

    int n=2;
    if(mp.find(2) != mp.end()){
        std::cout << n << " is present in map" << std::endl;
    }

    mp.erase(mp.begin());
    std::cout << "elements after deleting the first element: " << std::endl;
    std::cout << "key\t element" << std::endl;

    for(auto it=mp.begin(); it!=mp.end(); it++){
        std::cout << it->first << "\t" << it->second << std::endl;
    }

    std::cout << "the size of the map is: " << mp.size() << std::endl;

    if(mp.empty() == false)
        std::cout << "the map is not empty" << std::endl;
    else
        std::cout << "map is empty" << std::endl;
    
    mp.clear();
    std::cout << "size of map after clearing all elements: " << mp.size() << std::endl;

    return 0;
}