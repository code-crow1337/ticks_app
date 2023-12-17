import axios, { AxiosInstance } from "axios";
import useSWR from "swr";

const BASE_URL = "http://localhost:8080"; //fix to env file
let httpClient: AxiosInstance | undefined = undefined;
if (!httpClient) {
  httpClient = axios.create({
    baseURL: BASE_URL,
  });
}

export const useFetchApi = (path: string) => {
  const fetcher = (url: string) => httpClient?.get(url).then((res) => res.data);
  const { data, error } = useSWR(path, fetcher);

  return {
    data,
    isLoading: !error && !data,
    error,
  };
};
