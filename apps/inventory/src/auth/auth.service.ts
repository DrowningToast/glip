
import { PrismaClient } from "@prisma/client";

export class AuthService {
    constructor(
        private prisma: PrismaClient
    ) {}

    async checkUserExists(email: string) {
        const user = await this.prisma.user.findUnique({
            where: {
                email,
            },
        });

        return !!user;
    }

    async createUser({email, password}: {email: string, password: string}) {
        const hashedPassword = await Bun.password.hash(password, {
            algorithm: "bcrypt",
            cost: 10,
        });

        const user = await this.prisma.user.create({
            data: {
                email,
                password: hashedPassword,
            },
        });

        return user;
    }
}