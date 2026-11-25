import { Injectable } from '@nestjs/common';
import { PrismaService } from '../../infra/prisma.service';

export interface IAssetCreateRequest {
  id: string;
  symbol: string;
  price: number;
}
@Injectable()
export class AssetsCreateService {
  constructor(private readonly prismaService: PrismaService) {}

  async execute({ id, symbol, price }: IAssetCreateRequest) {
    return await this.prismaService.asset.create({
      data: {
        id: id,
        symbol: symbol,
        price: price,
      },
    });
  }
}
