#include <iostream>
#include <vector>
#include <algorithm>

int main() {
    int n, m;
    std::cin >> n >> m;

    std::vector<int> prices(n);
    for (int i = 0; i < n; ++i) {
        std::cin >> prices[i];
    }

    std::sort(prices.begin(), prices.end());

    int result = 0;
    int count = 0;

    for (int i = 0; i < n && count < m; ++i) {
        if (prices[i] < 0) {
            result += -prices[i];
            ++count;
        }
    }

    std::cout << result << std::endl;

    return 0;
}
