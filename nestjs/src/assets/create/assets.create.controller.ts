import { Body, Controller, Post } from '@nestjs/common';
import { AssetsCreateService } from './assets.create.service';

@Controller('assets')
export class AssetsCreateController {
  constructor(private readonly assetsCreateService: AssetsCreateService) {}

  @Post()
  public async create(@Body() body: {id: string;
  symbol: string;
    price: number
  }) {
    return await this.assetsCreateService.execute(body);
  }
}
