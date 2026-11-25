import { Injectable } from '@nestjs/common';
import { PrismaService } from '../../infra/prisma.service';

@Injectable()
export class WalletsGetAllService {
  constructor(private readonly prismaService: PrismaService) {}

  async execute() {
    return await this.prismaService.wallet.findMany();
  }
}
