import { Body, Controller, Get } from '@nestjs/common';
import { AssetsGetAllService } from './assets.get_all.service';

@Controller('assets')
export class AssetsGetAllController {
  constructor(private readonly assetsGetAllService: AssetsGetAllService) {}

  @Get()
  public async getAll() {
    return await this.assetsGetAllService.execute();
  }
}
