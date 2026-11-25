import { Injectable } from '@nestjs/common';
import { PrismaService } from '../../infra/prisma.service';

export interface IWalletCreateRequest {
  id: string;
}
@Injectable()
export class WalletsCreateService {
  constructor(private readonly prismaService: PrismaService) {}

  async execute({ id }: IWalletCreateRequest) {
    return await this.prismaService.wallet.create({
      data: {
        id: id,
      },
    });
  }
}
