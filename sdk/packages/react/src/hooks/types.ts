import type {
  DefaultError,
  QueryKey,
  UseQueryOptions,
} from '@tanstack/react-query'

export type QueryOpts<
  TQueryFnData = unknown,
  TError = DefaultError,
  TData = TQueryFnData,
  TQueryKey extends QueryKey = QueryKey,
> = Omit<
  UseQueryOptions<TQueryFnData, TError, TData, TQueryKey>,
  'queryKey' | 'queryFn' | 'enabled'
>
