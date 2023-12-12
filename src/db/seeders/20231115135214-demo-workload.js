'use strict'

const crypto = require('crypto')

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface, Sequelize) {
    const profileId = crypto.randomUUID()
    await queryInterface.bulkInsert('profiles', [
      {
        id: profileId,
        name: 'demo',
        description: 'This is an initial demo workload'
      }
    ])

    const answerId = crypto.randomUUID()
    await queryInterface.bulkInsert('profile-question-answer', [
      {
        id: answerId,
        name: 'demo',
        description: 'This is an initial demo answer'
      }
    ])

    const questionId = crypto.randomUUID()
    await queryInterface.bulkInsert('profile-question', [
      {
        id: questionId,
        name: 'demo',
        description: 'This is an initial demo question'
      }
    ])

    await queryInterface.bulkInsert('profile-question-answers', [
      {
        id: crypto.randomUUID(),
        questionId,
        answerId
      }
    ])

    await queryInterface.bulkInsert('profile-questions', [
      {
        id: crypto.randomUUID(),
        questionId: questionId,
        profileId: profileId
      }
    ])

    const pillarId = crypto.randomUUID()
    await queryInterface.bulkInsert('lens-pillar', [
      {
        id: pillarId,
        name: 'Operational Excellence',
        ref: 'ops_exec',
        description: ''
      }
    ])

    const pillarChoiceId1 = crypto.randomUUID()
    await queryInterface.bulkInsert('lens-pillar-choice', [
      {
        id: pillarChoiceId1,
        name: 'Choice 1',
        ref: 'choice_1',
        description: ''
      }
    ])

    const pillarChoiceId2 = crypto.randomUUID()
    await queryInterface.bulkInsert('lens-pillar-choice', [
      {
        id: pillarChoiceId2,
        name: 'Choice 2',
        ref: 'choice_2',
        description: ''
      }
    ])

    const pillarChoiceId3 = crypto.randomUUID()
    await queryInterface.bulkInsert('lens-pillar-choice', [
      {
        id: pillarChoiceId3,
        name: 'Choice 3',
        ref: 'choice_3',
        description: ''
      }
    ])

    const pillarChoiceId4 = crypto.randomUUID()
    await queryInterface.bulkInsert('lens-pillar-choice', [
      {
        id: pillarChoiceId4,
        name: 'Choice 4',
        ref: 'choice_4',
        description: ''
      }
    ])

    const pillarQuestionId = crypto.randomUUID()
    await queryInterface.bulkInsert('lens-pillar-question', [
      {
        id: pillarQuestionId,
        name: 'Question 1',
        ref: 'question_1',
        description:
          'Everyone needs to understand their part in enabling business success. Have shared goals in order to set priorities for resources. This will maximize the benefits of your efforts.'
      }
    ])

    await queryInterface.bulkInsert('lens-pillar-choices', [
      {
        id: crypto.randomUUID(),
        choiceId: pillarChoiceId1,
        questionId: pillarQuestionId
      },
      {
        id: crypto.randomUUID(),
        choiceId: pillarChoiceId2,
        questionId: pillarQuestionId
      },
      {
        id: crypto.randomUUID(),
        choiceId: pillarChoiceId3,
        questionId: pillarQuestionId
      },
      {
        id: crypto.randomUUID(),
        choiceId: pillarChoiceId4,
        questionId: pillarQuestionId
      }
    ])

    await queryInterface.bulkInsert('lens-pillar-questions', [
      {
        id: crypto.randomUUID(),
        questionId: pillarQuestionId,
        pillarId: pillarId
      }
    ])

    await queryInterface.bulkInsert('lens-pillar-question', [
      {
        id: crypto.randomUUID(),
        name: 'Question 1',
        ref: 'question_1',
        description: ''
      }
    ])

    const lensId = crypto.randomUUID()
    await queryInterface.bulkInsert('lenses', [
      {
        id: lensId,
        name: 'Web Application Security Lens',
        description: 'This is an initial demo lens',
        isDraft: false,
        spec: '{"version":1,"name":"SAP Lens","description":"SAP Lens","pillars":[{"id":"operational_excellence","name":"Operational Excellence","description":"Operational Excellence","questions":[{"id":"question_1","name":"Question 1","description":"Question 1","choices":[{"id":"choice_1","name":"Choice 1","description":"Choice 1"}],"risks":[{"risk":"HIGH","condition":"default"}]}]}]}'
      },
      {
        id: crypto.randomUUID(),
        name: 'SAP Workload',
        description: 'This is an initial demo lens',
        isDraft: true,
        spec: '{"version":1,"name":"SAP Lens","description":"SAP Lens","pillars":[{"id":"operational_excellence","name":"Operational Excellence","description":"Operational Excellence","questions":[{"id":"question_1","name":"Question 1","description":"Question 1","choices":[{"id":"choice_1","name":"Choice 1","description":"Choice 1"}],"risks":[{"risk":"HIGH","condition":"default"}]}]}]}'
      }
    ])

    await queryInterface.bulkInsert('lens-pillars', [
      {
        id: crypto.randomUUID(),
        lensId: lensId,
        pillarId: pillarId
      }
    ])

    const environmentId = crypto.randomUUID()
    await queryInterface.bulkInsert('environments', [
      {
        id: crypto.randomUUID(),
        name: 'production',
        description: 'Production environment'
      },
      {
        id: environmentId,
        name: 'staging',
        description: 'Staging environment'
      },
      {
        id: crypto.randomUUID(),
        name: 'development',
        description: 'Development environment'
      }
    ])

    const workloadId = crypto.randomUUID()
    await queryInterface.bulkInsert('workloads', [
      {
        id: workloadId,
        name: 'SAP Workload',
        description: 'SAP Workload',
        profilesId: profileId
      }
    ])

    await queryInterface.bulkInsert('workload-environment', [
      {
        id: crypto.randomUUID(),
        workloadId: workloadId,
        environmentId: environmentId
      }
    ])

    const workloadAnswerId = crypto.randomUUID()
    await queryInterface.bulkInsert('workload-lens-pillar-answer', [
      {
        id: workloadAnswerId,
        questionId: pillarQuestionId
      }
    ])

    await queryInterface.bulkInsert('workload-lens-pillar-answer-choices', [
      {
        id: crypto.randomUUID(),
        answerId: workloadAnswerId,
        choiceId: pillarChoiceId1
      },
      {
        id: crypto.randomUUID(),
        answerId: workloadAnswerId,
        choiceId: pillarChoiceId2
      }
    ])

    await queryInterface.bulkInsert('workload-lens-pillar-answers', [
      {
        id: crypto.randomUUID(),
        answerId: workloadAnswerId,
        workloadId: workloadId
      }
    ])

    await queryInterface.bulkInsert('workload-lens', [
      {
        id: crypto.randomUUID(),
        lensId: lensId,
        workloadId: workloadId
      }
    ])
  },

  async down(queryInterface, Sequelize) {
    await queryInterface.bulkDelete('workloads', null, {})
    await queryInterface.bulkDelete('profiles', null, {})
    await queryInterface.bulkDelete('lenses', null, {})
  }
}
