#include <bits/stdc++.h>

void charHashing();
void upperLowerHashing();
void bigNumberHashing() ;

int main() {
    bigNumberHashing();

    return 0;
}

void bigNumberHashing() {
    int n;
    std::cin >> n;

    int arr[n];
    for(int i=0; i<n; i++) {
        std::cin >> arr[i];
    }

    std::map<int, int> mp;
    for(int i=0; i<n; i++){
        mp[arr[i]]++;
    }

    int q;
    std::cin >> q;
    while(q--) {
        int number;
        std::cin >> number;
        std::cout << mp[number] << std::endl;
    }
}

void upperLowerHashing() {
    std::string s;
    std::cin >> s;

    int hash[256] = {0};
    for(int i=0; i<s.size(); i++) {
        hash[s[i]]++;
    }

    int q;
    std::cin >> q;

    while(q--) {
        char c;
        std::cin >> c;
        std::cout << hash[c] << std::endl;
    }
}

void charHashing() {
    std::string s;
    std::cin >> s;

    int hash[26] = {0};

    for(int i=0; i<s.size(); i++) {
        hash[s[i]-'a'] += 1; 
    }

    int q;
    std::cin >> q;
    while(q--) {
        char c;
        std::cin >> c;

        std::cout << hash[c-'a'] << std::endl;
    }
}



void intHashing() {
    int n;
    std::cin >> n;
    
    int arr[n];

    for(int i=0; i<n; i++) {
        std::cin >> arr[i];
    }

    int hash[13] = {0};
    for(int i=0; i<n; i++) {
        hash[arr[i]] += 1;
    }

    int q;
    std::cin >> q;
    while(q--) {
        int number;
        std::cin >> number;

        std::cout << hash[number] << std::endl;
    }
}