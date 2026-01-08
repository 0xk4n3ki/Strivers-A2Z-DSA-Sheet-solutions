#include <bits/stdc++.h>

int main() {
    std::multimap<int, int> mp;
    for(int i=1; i<=5; i++) {
        mp.insert({i, i*10});
    }
    mp.insert({4, 45});

    std::cout << "Elements present in the multimap: " << std::endl;
    std::cout << "Key\tElement" << std::endl;
    for(auto it = mp.begin(); it!=mp.end(); it++) {
        std::cout << it->first << "\t" << it->second << std::endl;
    }

    int n=2;
    if(mp.find(2) != mp.end()) {
        std::cout << n << " is present in multimap" << std::endl;
    }

    mp.erase(mp.begin());
    std::cout << "Elements after deleting the first element: " << std::endl;
    std::cout << "key\tElement" << std::endl;
    for(auto it=mp.begin(); it != mp.end(); it++) {
        std::cout << it->first << "\t" << it->second << std::endl;
    }

    std::cout << "The size of the multimap is: " << mp.size() << std::endl;
    
    if(mp.empty() == false)
        std::cout << "The multimap is not empty" << std::endl;
    else
        std::cout << "The multimap is empty" << std::endl;

    std::cout << "first element: " << mp.cbegin()->first << " " << mp.cbegin()->second << " " << mp.begin()->first << std::endl;

    mp.clear();
    std::cout << "Size of the multimap after clearing all the elements: " << mp.size() << std::endl;

    
    return 0;
}