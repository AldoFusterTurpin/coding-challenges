//ALDO FUSTER TURPIN
#include <iostream>
#include <vector>
using namespace std;


void printArray(const vector <int> &arr) {
    for (int i = 0; i < arr.size(); i++)
        cout << arr[i] << " ";
    cout << endl;
}


void swap(vector<int> &arr, const int &i, const int &j) {
    int tmp = arr[i];
    arr[i] = arr[j];
    arr[j] = tmp;
}


int partition(vector<int> &arr, int low, int high) {
    int pivot = arr[low];
    
    --low;
    ++high;
    
    while (low < high) {
        do ++low; while (arr[low] < pivot);
        do --high; while (arr[high] > pivot);
        if (low < high) swap(arr, low, high);
    }
    return high;
}
 


void quickSort(vector<int> &arr, const int &low, const int &high) {
    if (low < high) {
        int divisionPoint = partition(arr, low, high);
        quickSort(arr, low, divisionPoint);
        quickSort(arr, divisionPoint + 1, high);
    }
}


void sortArray(vector<int> &arr) {
    quickSort(arr, 0, arr.size()-1);
}


int main()
{
    int n;
    cin >> n;
    
    vector <int> arr(n);
    
    for (int i = 0; i < n; ++i)
        cin >> arr[i];
    
    sortArray(arr);
    
    printArray(arr);
    
    return 0;
}
