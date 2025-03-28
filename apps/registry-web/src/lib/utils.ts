import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
	return twMerge(clsx(inputs));
}

// Wrap the promise function to return in go convention
export function tryCatch<T>(
	fn: () => Promise<T>
): Promise<[Awaited<T> | undefined, Error | undefined]> {
	return Promise.resolve(fn())
		.then((res) => [res, undefined] satisfies [Awaited<T>, undefined])
		.catch((err) => [undefined, err] satisfies [undefined, Error]);
}
