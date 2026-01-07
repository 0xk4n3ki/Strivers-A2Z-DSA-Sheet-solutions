#include <bits/stdc++.h>

int main() {
    std::unordered_multimap<int, int> mp;

    for(int i=1; i<=5; i++) {
        mp.insert({i, i*10});
    }
    mp.insert({4, 35});

    std::cout << "elements present in unordered multimap: " << std::endl;
    std::cout << "key\t element" << std::endl;
    for(auto it=mp.begin(); it!=mp.end(); it++) {
        std::cout << it->first << "\t" << it->second << std::endl;
    } 

    int n=2;
    if(mp.find(n) != mp.end()) {
        std::cout << n << " is present in unordered multimap" << std::endl;
    }

    mp.erase(mp.begin());
    mp.erase(4);
    std::cout << "elements after deleting the firrst element:" << std::endl;
    std::cout << "key\t element" << std::endl;
    
    for(auto it=mp.begin(); it!=mp.end(); it++) {
        std::cout << it->first << "\t" << it->second << std::endl;
    }
    
    std::cout << "the size of the unordered multimap is: " << mp.size() << std::endl;;
    if(mp.empty() == false) {
        std::cout << "unordered multimap is not empty" << std::endl;
    }else {
        std::cout << "unordered multimap is empty" << std::endl;
    }

    mp.clear();
    std::cout << "size of the unordered multimap after clearing all the elements: " << mp.size() << std::endl;
}