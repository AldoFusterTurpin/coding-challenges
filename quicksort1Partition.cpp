//ALDO FUSTER TURPIN
#include <iostream>
#include <fstream>
#include <limits>
#include <vector>
#include <list>
#include <iterator>
using namespace std;


void copyToVec(const list<int> &l, vector<int> &v, int& i) {
    for (list<int>::const_iterator it = l.begin(); it != l.end(); ++it) {
        v[i] = *it;
        ++i;
    }
}


vector<int> quickSort(vector <int> arr) {
    list<int> left;
    list<int> equal;
    list<int> right;
    
    vector<int> result(arr.size());
    int index = 0;
    
    //pivot
    int p = arr[0];
    
    //construct left, equal and right
    for (int i = 0; i < arr.size(); i++){
        if (arr[i] < p) left.push_back(arr[i]);
        else if (arr[i] == p) equal.push_back(arr[i]);
        else right.push_back(arr[i]);
    }
    
    copyToVec(left, result, index);
    copyToVec(equal, result, index);
    copyToVec(right, result, index);
    
    return result;
}

int main()
{
    int n;
    cin >> n;
    
    vector <int> arr(n);
    
    for (int i = 0; i < n; ++i) {
        int element;
        cin >> element;
        arr[i] = element;
    }
    
    vector <int> result = quickSort(arr);
    for (int i = 0; i < result.size(); i++) {
        cout << result[i] << " ";
    }
    cout << endl;
}



