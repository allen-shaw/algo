__global__ void matmal(float* A, float* B, float* C, int m, int k, int n) {
    int row = blockIdx.y * blockDim.y + threadIdx.y;
    int col = blockIdx.x * blockDim.x + threadIdx.x;

    if (row < m and col < n) {
        float sum = 0.0f;
        for (int i = 0; i < k; i++) {
           sum += A[row * k + 1] * B[i * n + col];
        }
        C[row * n + col] = sum;
    }
}