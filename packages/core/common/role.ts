export const Role = {
	ROOT: "ROOT",
	WAREHOUSE: "WAREHOUSE",
	USER: "USER",
} as const;

export type Role = (typeof Role)[keyof typeof Role];
