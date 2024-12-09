export const fetchApi = async <T>(
  url: string,
  options: RequestInit = {}
): Promise<T> => {
  const { headers, ...restOptions } = options

  try {
    const response = await fetch(url, {
      ...restOptions,
      headers: {
        'Content-Type': 'application/json',
        ...headers,
      }
    })

    console.log("response", response)

    if (!response.ok) {
      const errorDetails = await response.text()
      return Promise.reject(errorDetails)
    }

    const data = (await response.json()) as T
    return data
  } catch (error) {
    throw error
  }
}