import { Injectable } from '@nestjs/common';
import { PrismaService } from '../../infra/prisma.service';

@Injectable()
export class AssetsGetAllService {
  constructor(private readonly prismaService: PrismaService) {}

  async execute() {
    return await this.prismaService.asset.findMany();
  }
}
