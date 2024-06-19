import { Module } from '@nestjs/common';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { TypeOrmModule } from '@nestjs/typeorm';

import { AppController } from './app.controller';
import { AppService } from './app.service';
import { HealthModule } from './health/health.module';
import { PersonModule } from './person/person.module';
import { V1Creation1718197500327 } from './migrations/1718197500327-V1_Creation';
import { V2NullablePersonLocation1718202607833 } from './migrations/1718202607833-V2_NullablePersonLocation';
import { V4DropPersonLocation1718203249447 } from './migrations/1718203249447-V4_DropPersonLocation';

@Module({
  imports: [
    ConfigModule.forRoot(),
    HealthModule,
    TypeOrmModule.forRootAsync({
      imports: [ConfigModule],
      useFactory: (configService: ConfigService) => ({
        type: 'postgres',
        host: configService.get('DB_HOST'),
        port: 5432,
        username: configService.get('DB_USER'),
        password: configService.get('DB_PASSWORD'),
        database: configService.get('DB_NAME'),
        autoLoadEntities: true,
        synchronize: false,
        migrations: [
          V1Creation1718197500327,
          V2NullablePersonLocation1718202607833,
          V4DropPersonLocation1718203249447,
        ],
        migrationsRun: true,
      }),
      inject: [ConfigService],
    }),
    PersonModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
