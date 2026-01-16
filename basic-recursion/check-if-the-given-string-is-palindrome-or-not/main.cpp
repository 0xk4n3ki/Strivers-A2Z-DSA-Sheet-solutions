#include <bits/stdc++.h>

bool checkPalindrome(std::string s, int left, int right) {
    if(left>=right) {
        return true;
    }
    if(!isalnum(static_cast<unsigned char>(s[left]))) {
        return checkPalindrome(s, left+1, right);
    }
    if(!isalnum(static_cast<unsigned char>(s[right]))) {
        return checkPalindrome(s, left, right-1);
    }
    if(tolower(static_cast<unsigned char>(s[left])) != tolower(static_cast<unsigned char>(s[right]))) {
        return false;
    }
    return checkPalindrome(s, left+1, right-1);
}

int main() {
    std::string s;
    std::cout << "Enter string: ";
    std::getline(std::cin, s);

    bool ans = checkPalindrome(s, 0, s.length()-1);
    if(ans) {
        std::cout << s << " is a palindrome" << std::endl;
    }else {
        std::cout << s << " is not a palindrome" << std::endl;
    }
    return 0;
}