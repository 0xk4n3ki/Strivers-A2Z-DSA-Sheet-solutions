#include <bits/stdc++.h>

void printStack(std::stack<int> s1) {
    std::stack<int> s2 = s1;
    while(!s2.empty()) {
        std::cout << s2.top() << " ";
        s2.pop();
    }
    std::cout << std::endl;
}

int main() {
    std::stack<int> s;
    for(int i=0; i<=5; i++) {
        s.push(i);
    }

    std::cout << "The elements of the stack are: ";
    printStack(s);

    std::cout << "The size of the stack: " << s.size() << std::endl;
    std::cout << "The top element of the stack: " << s.top() << std::endl;
    std::cout << "Pop the top element: " << std::endl;
    s.pop();

    printStack(s);

    return 0;
}