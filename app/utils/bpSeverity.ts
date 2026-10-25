export function getBloodPressureSeverity(systolic: number, diastolic: number) {
  if (systolic <= 120 && diastolic <= 80) {
    return {
      text: 'Normal',
      class: 'bg-green-100 text-green-800'
    }
  } else if ((systolic > 120 && systolic <= 139) || (diastolic > 80 && diastolic <= 89)) {
    return {
      text: 'Pre Hypertension',
      class: 'bg-yellow-100 text-yellow-800'
    }
  } else if ((systolic >= 140 && systolic <= 159) || (diastolic >= 90 && diastolic <= 99)) {
    return {
      text: 'Stage 1 Hypertension',
      class: 'bg-red-100 text-red-800'
    }
  } else if (systolic >= 160 && diastolic >= 100) {
    return {
      text: 'Stage 2 Hypertension',
      class: 'bg-red-200 text-red-900'
    }
  } else {
    return {
      text: 'N/A',
      class: 'bg-gray-100 text-gray-800'
    }
  }
}
