<?php

namespace Database\Factories;

use App\BloodPressure;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\BloodPressure>
 */
class BloodPressureFactory extends Factory
{
    /**
     * The name of the factory's corresponding model.
     *
     * @var string
     */
    protected $model = BloodPressure::class;

    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            'user_id' => 0,
            'systolic' => fake()->numberBetween(100, 180),
            'diastolic' => fake()->numberBetween(60, 120),
            'pulse' => fake()->numberBetween(60, 90),
            'reading_date' => now(),
        ];
    }
}
